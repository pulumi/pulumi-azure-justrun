# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = ['WebappArgs', 'Webapp']

@pulumi.input_type
class WebappArgs:
    def __init__(__self__, *,
                 app_sku_name: Optional[pulumi.Input[str]] = None,
                 app_sku_tier: Optional[pulumi.Input[str]] = None,
                 container_public_access: Optional[pulumi.Input['PublicAccess']] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_sku_name: Optional[pulumi.Input['StorageSkuName']] = None):
        """
        The set of arguments for constructing a Webapp resource.
        :param pulumi.Input[str] app_sku_name: The tier of the compute instance running the server. Also see appSkuName
        :param pulumi.Input[str] app_sku_tier: The name of the compute instance running the server. Also see appSkuTier
        :param pulumi.Input['PublicAccess'] container_public_access: The public access level of the BlobContainer containg the website data.
        :param pulumi.Input[str] file_path: The relative file path to the folder containing web files.
        :param pulumi.Input[str] name_prefix: The name prefix given to child resources of this component. Should not contain dashes.
        :param pulumi.Input[str] resource_group_name: The resource group to use. One will be created if not provided.
        :param pulumi.Input[str] storage_account_name: The name of the storage account to use. One will be created if not provided.
        :param pulumi.Input['StorageSkuName'] storage_sku_name: The SKU name of the storage account created, if storageAccount is not provided
        """
        if app_sku_name is not None:
            pulumi.set(__self__, "app_sku_name", app_sku_name)
        if app_sku_tier is not None:
            pulumi.set(__self__, "app_sku_tier", app_sku_tier)
        if container_public_access is not None:
            pulumi.set(__self__, "container_public_access", container_public_access)
        if file_path is None:
            file_path = './www'
        if file_path is not None:
            pulumi.set(__self__, "file_path", file_path)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if resource_group_name is not None:
            pulumi.set(__self__, "resource_group_name", resource_group_name)
        if storage_account_name is not None:
            pulumi.set(__self__, "storage_account_name", storage_account_name)
        if storage_sku_name is not None:
            pulumi.set(__self__, "storage_sku_name", storage_sku_name)

    @property
    @pulumi.getter(name="appSkuName")
    def app_sku_name(self) -> Optional[pulumi.Input[str]]:
        """
        The tier of the compute instance running the server. Also see appSkuName
        """
        return pulumi.get(self, "app_sku_name")

    @app_sku_name.setter
    def app_sku_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_sku_name", value)

    @property
    @pulumi.getter(name="appSkuTier")
    def app_sku_tier(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the compute instance running the server. Also see appSkuTier
        """
        return pulumi.get(self, "app_sku_tier")

    @app_sku_tier.setter
    def app_sku_tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_sku_tier", value)

    @property
    @pulumi.getter(name="containerPublicAccess")
    def container_public_access(self) -> Optional[pulumi.Input['PublicAccess']]:
        """
        The public access level of the BlobContainer containg the website data.
        """
        return pulumi.get(self, "container_public_access")

    @container_public_access.setter
    def container_public_access(self, value: Optional[pulumi.Input['PublicAccess']]):
        pulumi.set(self, "container_public_access", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> Optional[pulumi.Input[str]]:
        """
        The relative file path to the folder containing web files.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_path", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The name prefix given to child resources of this component. Should not contain dashes.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource group to use. One will be created if not provided.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="storageAccountName")
    def storage_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the storage account to use. One will be created if not provided.
        """
        return pulumi.get(self, "storage_account_name")

    @storage_account_name.setter
    def storage_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_name", value)

    @property
    @pulumi.getter(name="storageSkuName")
    def storage_sku_name(self) -> Optional[pulumi.Input['StorageSkuName']]:
        """
        The SKU name of the storage account created, if storageAccount is not provided
        """
        return pulumi.get(self, "storage_sku_name")

    @storage_sku_name.setter
    def storage_sku_name(self, value: Optional[pulumi.Input['StorageSkuName']]):
        pulumi.set(self, "storage_sku_name", value)


class Webapp(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_sku_name: Optional[pulumi.Input[str]] = None,
                 app_sku_tier: Optional[pulumi.Input[str]] = None,
                 container_public_access: Optional[pulumi.Input['PublicAccess']] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_sku_name: Optional[pulumi.Input['StorageSkuName']] = None,
                 __props__=None):
        """
        This represents a web app component resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_sku_name: The tier of the compute instance running the server. Also see appSkuName
        :param pulumi.Input[str] app_sku_tier: The name of the compute instance running the server. Also see appSkuTier
        :param pulumi.Input['PublicAccess'] container_public_access: The public access level of the BlobContainer containg the website data.
        :param pulumi.Input[str] file_path: The relative file path to the folder containing web files.
        :param pulumi.Input[str] name_prefix: The name prefix given to child resources of this component. Should not contain dashes.
        :param pulumi.Input[str] resource_group_name: The resource group to use. One will be created if not provided.
        :param pulumi.Input[str] storage_account_name: The name of the storage account to use. One will be created if not provided.
        :param pulumi.Input['StorageSkuName'] storage_sku_name: The SKU name of the storage account created, if storageAccount is not provided
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WebappArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This represents a web app component resource

        :param str resource_name: The name of the resource.
        :param WebappArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebappArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_sku_name: Optional[pulumi.Input[str]] = None,
                 app_sku_tier: Optional[pulumi.Input[str]] = None,
                 container_public_access: Optional[pulumi.Input['PublicAccess']] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 storage_account_name: Optional[pulumi.Input[str]] = None,
                 storage_sku_name: Optional[pulumi.Input['StorageSkuName']] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebappArgs.__new__(WebappArgs)

            __props__.__dict__["app_sku_name"] = app_sku_name
            __props__.__dict__["app_sku_tier"] = app_sku_tier
            __props__.__dict__["container_public_access"] = container_public_access
            if file_path is None:
                file_path = './www'
            __props__.__dict__["file_path"] = file_path
            __props__.__dict__["name_prefix"] = name_prefix
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["storage_account_name"] = storage_account_name
            __props__.__dict__["storage_sku_name"] = storage_sku_name
            __props__.__dict__["url"] = None
        super(Webapp, __self__).__init__(
            'azure-justrun:index:webapp',
            resource_name,
            __props__,
            opts,
            remote=True)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL of the web app
        """
        return pulumi.get(self, "url")

