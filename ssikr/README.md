

## 실습 준비사항
A. Go언어
              1) Go언어의 설치
                  https://go.dev/doc/install
              2) goland 설치(무료버전)
                  https://www.jetbrains.com/ko-kr/go/download/
              3) Go언어의 기본 사용법
                  https://go-tour-ko.appspot.com/welcome/1
              4) goland에서 기본 예제 작성 후 실행해 보기

	B. protoc 설치
             아래 링크의 중간 install 부분을 참고하시거나 검색을 통해 해당 노트북에 맞게 설치 부탁드립니다.
             https://velog.io/@milkcoke/Go-grpc-%EC%82%AC%EC%9A%A9%ED%95%B4%EB%B3%B4%EA%B8%B0

		> brew install protobuf

	C. WSL설치(Windows 환경의 경우만)
              https://docs.microsoft.com/ko-kr/windows/wsl/install

	D. protobuf
             (이 내용은 한 번 보시고 오시면 실습하실 때 도움이 되실 것입니다.)
             https://developers.google.com/protocol-buffers
             https://corgipan.tistory.com/8
             https://blog.naver.com/alice_k106/221617347519

	E. BloomRPC 설치
	    - https://github.com/bloomrpc/bloomrpc

		> brew install --cask bloomrpc

	F. Docker
	    - https://www.docker.com/get-started/

### proto 파일을 만든뒤

protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    protos/**/*.proto


기본 홀더 시나리오


