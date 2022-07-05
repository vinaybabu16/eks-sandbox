package main

import (
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an EKS cluster with the default configuration.
		cluster, err := eks.NewCluster(ctx, "argocd-playground", &eks.ClusterArgs{
			SubnetIds: pulumi.StringArray{
				pulumi.String("subnet-05fd0b3b5d5e326d0"),
				pulumi.String("subnet-02b6f54eafb123a71"),
			},
			MaxSize:         pulumi.IntPtr(5),
			MinSize:         pulumi.IntPtr(5),
			DesiredCapacity: pulumi.IntPtr(5),
		})
		if err != nil {
			return err
		}

		// Export the cluster's kubeconfig.
		ctx.Export("kubeconfig", cluster.Kubeconfig)
		return nil
	})
}
