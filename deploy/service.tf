resource "google_cloud_run_v2_service" "service" {
  name     = var.service_name
  location = var.gcp_region

  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = var.service_image

      env {
        name  = "SYNOPTIC_API_TOKEN"
        value = var.synoptic_api_token
      }
    }

    scaling {
      max_instance_count = var.max_instance_count
      min_instance_count = var.min_instance_count
    }
  }

  depends_on = [google_project_service.run_api]
}

resource "google_project_service" "run_api" {
  service = "run.googleapis.com"

  disable_on_destroy = true
}
