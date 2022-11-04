# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
import pulumi_aws

__all__ = ['NodeGroupArgs', 'NodeGroup']

@pulumi.input_type
class NodeGroupArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 name: pulumi.Input[str],
                 instance_type: Optional[pulumi.Input[str]] = None,
                 max_count: Optional[pulumi.Input[float]] = None,
                 min_count: Optional[pulumi.Input[float]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a NodeGroup resource.
        :param pulumi.Input[str] cluster_name: The ECS cluster name.
        :param pulumi.Input[str] name: The name that identies the resource.
        :param pulumi.Input[str] instance_type: The aws instance type to use for the nodes. Defaults to "t3.medium".
        :param pulumi.Input[float] max_count: The maxium number of nodes running in the node group. Defaults to 2.
        :param pulumi.Input[float] min_count: The minimum number of nodes running in the node group. Defaults to 1.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: The list of subnets ids where the nodes will be deployed.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "name", name)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if max_count is not None:
            pulumi.set(__self__, "max_count", max_count)
        if min_count is not None:
            pulumi.set(__self__, "min_count", min_count)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        The ECS cluster name.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name that identies the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The aws instance type to use for the nodes. Defaults to "t3.medium".
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[pulumi.Input[float]]:
        """
        The maxium number of nodes running in the node group. Defaults to 2.
        """
        return pulumi.get(self, "max_count")

    @max_count.setter
    def max_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_count", value)

    @property
    @pulumi.getter(name="minCount")
    def min_count(self) -> Optional[pulumi.Input[float]]:
        """
        The minimum number of nodes running in the node group. Defaults to 1.
        """
        return pulumi.get(self, "min_count")

    @min_count.setter
    def min_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "min_count", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of subnets ids where the nodes will be deployed.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)


class NodeGroup(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 max_count: Optional[pulumi.Input[float]] = None,
                 min_count: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        NodeGroup is a component that deploy a Node Group for a Kubernetes cluster.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: The ECS cluster name.
        :param pulumi.Input[str] instance_type: The aws instance type to use for the nodes. Defaults to "t3.medium".
        :param pulumi.Input[float] max_count: The maxium number of nodes running in the node group. Defaults to 2.
        :param pulumi.Input[float] min_count: The minimum number of nodes running in the node group. Defaults to 1.
        :param pulumi.Input[str] name: The name that identies the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: The list of subnets ids where the nodes will be deployed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NodeGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        NodeGroup is a component that deploy a Node Group for a Kubernetes cluster.

        :param str resource_name: The name of the resource.
        :param NodeGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodeGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 max_count: Optional[pulumi.Input[float]] = None,
                 min_count: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NodeGroupArgs.__new__(NodeGroupArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["max_count"] = max_count
            __props__.__dict__["min_count"] = min_count
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["subnet_ids"] = subnet_ids
            __props__.__dict__["launch_template"] = None
            __props__.__dict__["node_group"] = None
            __props__.__dict__["role"] = None
            __props__.__dict__["role_policy_attachments"] = None
        super(NodeGroup, __self__).__init__(
            'cloud-toolkit-aws:kubernetes:NodeGroup',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter(name="launchTemplate")
    def launch_template(self) -> pulumi.Output['pulumi_aws.ec2.LaunchTemplate']:
        """
        The EC2 Launch Template used to provision nodes.
        """
        return pulumi.get(self, "launch_template")

    @property
    @pulumi.getter(name="nodeGroup")
    def node_group(self) -> pulumi.Output['pulumi_aws.eks.NodeGroup']:
        """
        The EKS Node Group.
        """
        return pulumi.get(self, "node_group")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output['pulumi_aws.iam.Role']:
        """
        The IAM Role assumed by the EKS Nodes.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="rolePolicyAttachments")
    def role_policy_attachments(self) -> pulumi.Output[Sequence['pulumi_aws.iam.RolePolicyAttachment']]:
        """
        The list of IAM Role Policy Attachment used to attach IAM Roles to the EKS Node Group.
        """
        return pulumi.get(self, "role_policy_attachments")

