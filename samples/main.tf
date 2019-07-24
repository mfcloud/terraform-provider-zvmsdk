provider "zvm" {
  alias = "s8080"
  uri = "http://9.60.29.50:8080"
}

provider "zvm" {
  alias = "s8081"
  uri   = "http://127.0.0.1:8080"
}
  
resource "zvm_guest" "test1" {
  provider = zvm.s8080
  userid = "TERRTST1"
  userprofile = "osdflt"
  imageid = "5e0cdd2b-5f8b-4e38-b811-072a5073e89f"

  disklist {
       size = "4g"
       diskpool = "ECKD:xcateckd"
  }
}

resource "zvm_interface" "domain-1-if" {
  provider = zvm.s8080
  userid = zvm_guest.test1.id
  osversion = "rhel7.6"

  networks {
       ip = "1.2.3.4"
       vdev = "1000"
       cidr = "1.2.3.0/24"
  }
  active = "0"
}
