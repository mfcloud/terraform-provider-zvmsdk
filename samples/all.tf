provider "zvm" {
  alias = "s8080"
  uri = "http://127.0.0.1:8080"
}

provider "zvm" {
  alias = "s8081"
  uri   = "http://127.0.0.1:8081"
}
  
resource "zvm_image" "image1" {
  provider = zvm.s8080
  name = "image1"
  meta = {
    osversion = "rhel7.6"
    md5sum = "12345678"
  }
  url = "file://a/b/c"
  remotehost = "remote@1.2.3.4"
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
  }

  disklist {
       size = "2g"
       boot = 1
       format = "ext4"
  }
}

resource "zvm_interface" "domain-1-if" {
  provider = zvm.s8080
  userid = "domain-1"
  osversion = "rhel7.6"

  networks {
       ip = "1.2.3.4"
       vdev = "1000"
  }

  networks {
       ip = "5.6.7.8"
       vdev = "2000"
       gateway = "6.7.8.9"
  }
}

resource "zvm_vswitchgrant" "couple1" {
  provider = zvm.s8080
  nic = "1000"
  userid = "domain-1"
  vswitch = "vsw1"
}

terraform {
  required_version = ">= 0.12"
}
