# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'PublicAccess',
    'SkuName',
]


class PublicAccess(str, Enum):
    """
    Duplicates azure-native:storage:PublicAccess
    """
    CONTAINER = "Container"
    BLOB = "Blob"
    NONE = "None"


class SkuName(str, Enum):
    """
    Duplicates azure-native:storage:SkuName
    """
    STANDARD_LRS = "Standard_LRS"
    STANDARD_GRS = "Standard_GRS"
    STANDARD_RAGRS = "Standard_RAGRS"
    STANDARD_ZRS = "Standard_ZRS"
    PREMIUM_LRS = "Premium_LRS"
    PREMIUM_ZRS = "Premium_ZRS"
    STANDARD_GZRS = "Standard_GZRS"
    STANDARD_RAGZRS = "Standard_RAGZRS"
