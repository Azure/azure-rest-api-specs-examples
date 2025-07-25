from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.resources import ResourceManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python calculate_template_hash.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ResourceManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.deployments.calculate_template_hash(
        template={
            "$schema": "http://schemas.management.azure.com/deploymentTemplate?api-version=2014-04-01-preview",
            "contentVersion": "1.0.0.0",
            "outputs": {"string": {"type": "string", "value": "myvalue"}},
            "parameters": {"string": {"type": "string"}},
            "resources": [],
            "variables": {
                "array": [1, 2, 3, 4],
                "bool": True,
                "int": 42,
                "object": {"object": {"location": "West US", "vmSize": "Large"}},
                "string": "string",
            },
        },
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/deployments/stable/2022-09-01/examples/CalculateTemplateHash.json
if __name__ == "__main__":
    main()
