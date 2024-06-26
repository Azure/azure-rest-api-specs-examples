from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.healthcareapis import HealthcareApisManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-healthcareapis
# USAGE
    python fhir_services_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HealthcareApisManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.fhir_services.begin_create_or_update(
        resource_group_name="testRG",
        workspace_name="workspace1",
        fhir_service_name="fhirservice1",
        fhirservice={
            "identity": {
                "type": "UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/subid/resourcegroups/testRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-mi": {}
                },
            },
            "kind": "fhir-R4",
            "location": "westus",
            "properties": {
                "acrConfiguration": {"loginServers": ["test1.azurecr.io"]},
                "authenticationConfiguration": {
                    "audience": "https://azurehealthcareapis.com",
                    "authority": "https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc",
                    "smartIdentityProviders": [
                        {
                            "applications": [
                                {
                                    "allowedDataActions": ["Read"],
                                    "audience": "22222222-2222-2222-2222-222222222222",
                                    "clientId": "22222222-2222-2222-2222-222222222222",
                                }
                            ],
                            "authority": "https://login.b2clogin.com/11111111-1111-1111-1111-111111111111/v2.0",
                        }
                    ],
                    "smartProxyEnabled": True,
                },
                "corsConfiguration": {
                    "allowCredentials": False,
                    "headers": ["*"],
                    "maxAge": 1440,
                    "methods": ["DELETE", "GET", "OPTIONS", "PATCH", "POST", "PUT"],
                    "origins": ["*"],
                },
                "encryption": {
                    "customerManagedKeyEncryption": {
                        "keyEncryptionKeyUrl": "https://mykeyvault.vault.azure.net/keys/myEncryptionKey/myKeyVersion"
                    }
                },
                "exportConfiguration": {"storageAccountName": "existingStorageAccount"},
                "implementationGuidesConfiguration": {"usCoreMissingData": False},
                "importConfiguration": {
                    "enabled": False,
                    "initialImportMode": False,
                    "integrationDataStore": "existingStorageAccount",
                },
            },
            "tags": {"additionalProp1": "string", "additionalProp2": "string", "additionalProp3": "string"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/fhirservices/FhirServices_Create.json
if __name__ == "__main__":
    main()
