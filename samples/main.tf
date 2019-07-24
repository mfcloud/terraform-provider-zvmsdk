provider "zvm" {
  alias = "s8080"
  uri = "http://9.60.29.50:8080"
}

provider "zvm" {
  alias = "s8081"
  uri   = "http://127.0.0.1:8080"
}
  
resource "zvm_guest" "guest1" {
  provider = zvm.s8080
  userid = "TERRTST1"
  userprofile = "osdflt"
  imageid = "5e0cdd2b-5f8b-4e38-b811-072a5073e89f"

  disklist {
       size = "4g"
       diskpool = "ECKD:xcateckd"
  }
}
