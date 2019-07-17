provider "zvm" {
  alias = "s8080"
  uri = "http://127.0.0.1:8080"
}

provider "zvm" {
  alias = "s8081"
  uri   = "http://127.0.0.1:8081"
}
  

resource "zvm_guest" "guest1" {
  provider = zvm.s8080
  userid = "domain-1"
  userprofile = "p1"
  imageid = "image1"
  diskpool = "ECKD:eckdpool1"

  disklist {
       size = "1g"
       boot = 0
       format = "ext4"
  }

  disklist {
       size = "2g"
       boot = 1
       format = "ext4"
  }
}

terraform {
  required_version = ">= 0.12"
}
