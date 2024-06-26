from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.healthcareapis import HealthcareApisManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-healthcareapis
# USAGE
    python service_create_in_data_sovereign_region_with_cmk_enabled.py

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

    response = client.services.begin_create_or_update(
        resource_group_name="rg1",
        resource_name="service1",
        service_description={
            "identity": {"type": "SystemAssigned"},
            "kind": "fhir-R4",
            "location": "Southeast Asia",
            "properties": {
                "accessPolicies": [
                    {"objectId": "c487e7d1-3210-41a3-8ccc-e9372b78da47"},
                    {"objectId": "5b307da8-43d4-492b-8b66-b0294ade872f"},
                ],
                "authenticationConfiguration": {
                    "audience": "https://azurehealthcareapis.com",
                    "authority": "https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc",
                    "smartProxyEnabled": True,
                },
                "corsConfiguration": {
                    "allowCredentials": False,
                    "headers": ["*"],
                    "maxAge": 1440,
                    "methods": ["DELETE", "GET", "OPTIONS", "PATCH", "POST", "PUT"],
                    "origins": ["*"],
                },
                "cosmosDbConfiguration": {
                    "crossTenantCmkApplicationId": "de3fbeef-8c3a-428e-8b9f-4d229c8a85f4",
                    "keyVaultKeyUri": "https://my-vault.vault.azure.net/keys/my-key",
                    "offerThroughput": 1000,
                },
                "exportConfiguration": {"storageAccountName": "existingStorageAccount"},
                "privateEndpointConnections": [],
                "publicNetworkAccess": "Disabled",
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/ServiceCreateInDataSovereignRegionWithCmkEnabled.json
if __name__ == "__main__":
    main()
