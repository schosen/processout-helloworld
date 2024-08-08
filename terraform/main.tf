
resource "kubernetes_namespace" "checkout" {
  metadata {
    name = "checkout"
  }
}

resource "kubernetes_deployment" "processout-helloworld" {
  metadata {
    name      = "processout-helloworld"
    namespace = kubernetes_namespace.checkout.metadata[0].name
    labels = {
      app = "processout-helloworld"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "processout-helloworld"
      }
    }

    template {
      metadata {
        labels = {
          app = "processout-helloworld"
        }
      }

      spec {
        container {
          name  = "processout-helloworld"
          image = "${var.docker_username}/processout-helloworld:latest"

          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "processout-helloworld" {
  metadata {
    name      = "processout-helloworld-service"
    namespace = kubernetes_namespace.checkout.metadata[0].name
  }

  spec {
    selector = {
      app = kubernetes_deployment.helloworld.spec[0].template[0].metadata[0].labels.app
    }

    port {
      port        = 80
      target_port = 8080
    }

    type = "NodePort"
  }
}
