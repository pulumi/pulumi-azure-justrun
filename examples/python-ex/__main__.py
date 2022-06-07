"""An Azure RM Python Pulumi program"""

import pulumi
import pulumi_azure_justrun 


webapp = pulumi_azure_justrun.Webapp("mywebapp", file_path="./www")

pulumi.export("url",webapp.url)
