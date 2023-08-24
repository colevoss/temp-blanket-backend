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

variable "max_instance_count" {
  type = number
}

variable "min_instance_count" {
  type = number
}

variable "synoptic_api_token" {
  type = string
}
