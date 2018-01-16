resource "google_compute_instance_group" "router-lb-0" {
  name        = "${var.env_id}-router-lb-0-${element(var.zones, 0)}"
  description = "terraform generated instance group that is multi-zone for https loadbalancing"
  zone        = "${element(var.zones, 0)}"

  named_port {
    name = "https"
    port = "443"
  }
}

resource "google_compute_instance_group" "router-lb-1" {
  name        = "${var.env_id}-router-lb-1-${element(var.zones, 1)}"
  description = "terraform generated instance group that is multi-zone for https loadbalancing"
  zone        = "${element(var.zones, 1)}"

  named_port {
    name = "https"
    port = "443"
  }
}

resource "google_compute_instance_group" "router-lb-2" {
  count       = "${var.restrict_instance_groups ? 0 : 1}"
  name        = "${var.env_id}-router-lb-2-${element(var.zones, 2)}"
  description = "terraform generated instance group that is multi-zone for https loadbalancing"
  zone        = "${element(var.zones, 2)}"

  named_port {
    name = "https"
    port = "443"
  }
}
