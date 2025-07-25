from azure.identity import DefaultAzureCredential

from azure.mgmt.hybridcompute import HybridComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcompute
# USAGE
    python license_profile_update.py

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

    response = client.license_profiles.begin_update(
        resource_group_name="myResourceGroup",
        machine_name="myMachine",
        parameters={
            "properties": {
                "esuProfile": {"assignedLicense": "{LicenseResourceId}"},
                "productProfile": {
                    "productFeatures": [{"name": "Hotpatch", "subscriptionStatus": "Enable"}],
                    "productType": "WindowsServer",
                    "subscriptionStatus": "Enable",
                },
                "softwareAssurance": {"softwareAssuranceCustomer": True},
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/licenseProfile/LicenseProfile_Update.json
if __name__ == "__main__":
    main()
