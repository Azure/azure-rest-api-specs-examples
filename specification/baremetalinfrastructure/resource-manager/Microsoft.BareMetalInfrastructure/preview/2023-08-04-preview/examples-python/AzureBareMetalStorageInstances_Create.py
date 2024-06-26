from azure.identity import DefaultAzureCredential
from azure.mgmt.baremetalinfrastructure import BareMetalInfrastructureClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-baremetalinfrastructure
# USAGE
    python azure_bare_metal_storage_instances_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = BareMetalInfrastructureClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.azure_bare_metal_storage_instances.create(
        resource_group_name="myResourceGroup",
        azure_bare_metal_storage_instance_name="myAzureBareMetalStorageInstance",
        request_body_parameters={
            "location": "westus2",
            "properties": {
                "azureBareMetalStorageInstanceUniqueIdentifier": "23415635-4d7e-41dc-9598-8194f22c24e9",
                "storageProperties": {
                    "generation": "Gen4",
                    "hardwareType": "NetApp",
                    "offeringType": "EPIC",
                    "provisioningState": "Succeeded",
                    "storageBillingProperties": {"azureBareMetalStorageInstanceSize": "", "billingMode": "PAYG"},
                    "storageType": "FC",
                    "workloadType": "ODB",
                },
            },
            "tags": {"key": "value"},
        },
    )
    print(response)


# x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalStorageInstances_Create.json
if __name__ == "__main__":
    main()
