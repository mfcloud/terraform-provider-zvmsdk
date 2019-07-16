provider "zvm" {
  alias = "s8080"
  uri = "http://127.0.0.1:8080"
}

provider "zvm" {
  alias = "s8081"
  uri   = "http://127.0.0.1:8081"
}
  

resource "zvm_guest" "guest1" {
  provider = zvm.s8081
  userid = "domain-1"
  userprofile = "p1"
  imageid = "image1"
  diskpool = "ECKD:eckdpool1"

  disklist {
       size = "1g"
       boot = true
       diskpool = "ECKD:eckdpool1"
  }

  disklist {
       size = "2g"
       boot = false
       diskpool = "ECKD:eckdpool1"
  }
}

terraform {
  required_version = ">= 0.12"
}
