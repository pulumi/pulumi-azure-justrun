"""An Azure RM Python Pulumi program"""

import pulumi
import pulumi_azure_justrun 


containerapp = pulumi_azure_justrun.Containerapp("containerapp", image_directory="node-app")

pulumi.export("url",containerapp.url)
