variable "gcp_project_id" {
  type = string
}

variable "gcp_region" {
  type = string
}

variable "service_image" {
  type = string
}

variable "service_name" {
  type        = string
  description = "Name of the Cloud Run service"
}
