package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func createTar() {
	fw, err := os.Create("/Users/daniel/Desktop/demo.tar")
	handleError(err)
	fmt.Println(fw.Name())
	defer fw.Close()
	tw := tar.NewWriter(fw)
	defer tw.Close()
	fr, err := os.Open("/Users/daniel/Desktop/demo.txt")
	handleError(err)
	fmt.Println(fr.Name())
	defer fr.Close()

	fi, err := fr.Stat()
	handleError(err)
	fmt.Println(fi.Name())
	hdr := new(tar.Header)
	hdr.Name = fi.Name()
	hdr.Size = fi.Size()
	hdr.Mode = int64(fi.Mode())
	hdr.ModTime = fi.ModTime()
	err = tw.WriteHeader(hdr)
	handleError(err)
	_, err = io.Copy(tw, fr)
	handleError(err)
	//os.Remove(fw.Name())

}
func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func handleTar()  {
		fr, err := os.Open("/Users/daniel/Desktop/demo.tar")	// 打开tar包文件，返回*io.Reader
		handleError(err)	// handleError为错误处理函数，下同
		defer fr.Close()

		// 实例化新的tar.Reader
		tr := tar.NewReader(fr)

		for	{
			hdr, err := tr.Next()	// 获取下一个文件，第一个文件也用此方法获取
			if err == io.EOF {
				break	// 已读到文件尾
			}
			handleError(err)

			// 通过创建文件获得*io.Writer
			fw, _ := os.Create("demo/" + hdr.Name)
			handleError(err)

			// 拷贝数据
			_, err = io.Copy(fw, tr)
			handleError(err)
		}
}
/*创建zip*/
func createZip()  {
	fw,err:=os.Create("/Users/daniel/Desktop/demo.zip")
	handleError(err)
	defer fw.Close()
	// 实例化新的zip.Writer
	zw := zip.NewWriter(fw)
	defer zw.Close()
	fr,err:=os.Open("/Users/daniel/Desktop/demo.txt")
	handleError(err)
	defer fr.Close()

	// 获取文件信息
	fi, err := fr.Stat()
	handleError(err)

	// 创建zip.FileHeader
	fh := new(zip.FileHeader)
	fh.Name = fi.Name()
	fh.UncompressedSize64=uint64(fi.Size())
	//fh.UncompressedSize = uint32(fi.Size())
	fw1, err := zw.CreateHeader(fh)
	handleError(err)

	// 读取文件数据
	buf := make([]byte, fi.Size())
	_, err = fr.Read(buf)
	handleError(err)

	// 写入数据到zip包
	_, err = fw1.Write(buf)
	handleError(err)
	fmt.Println("create zip ")

}
/*处理zip*/
func handleZip()  {
	// 用zip.OpenReader打开zip包文件
	rc, err := zip.OpenReader("/Users/daniel/Desktop/demo.zip")
	handleError(err)	// handleError为错误处理函数，下同		defer rc.Close()

	// 使用循环逐个读取zip包内的单独文件
	for _, f := range rc.File {
		// 打开包中的文件
		r, err := f.Open()
		handleError(err)

		// 将包中的文件数据写出
		fw, _ := os.Create(f.FileInfo().Name())
		handleError(err)
		_, err = io.Copy(fw, r)
		handleError(err)
		fmt.Println(len(fw.Name()))
	}

}
func main() {
	createTar()
	fmt.Println("get started")
	handleTar()
	createZip()
	handleZip()
}

