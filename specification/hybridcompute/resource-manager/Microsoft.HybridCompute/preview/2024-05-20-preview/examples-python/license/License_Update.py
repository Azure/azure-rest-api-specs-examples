from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.hybridcompute import HybridComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcompute
# USAGE
    python license_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscriptionId}",
    )

    response = client.licenses.begin_update(
        resource_group_name="myResourceGroup",
        license_name="{licenseName}",
        parameters={
            "properties": {
                "licenseDetails": {
                    "edition": "Datacenter",
                    "processors": 6,
                    "state": "Activated",
                    "target": "Windows Server 2012",
                    "type": "pCore",
                },
                "licenseType": "ESU",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/license/License_Update.json
if __name__ == "__main__":
    main()
