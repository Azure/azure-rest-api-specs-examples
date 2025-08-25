from azure.identity import DefaultAzureCredential

from azure.mgmt.security import SecurityCenter

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-security
# USAGE
    python put_advanced_threat_protection_settings_example.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityCenter(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.advanced_threat_protection.create(
        resource_id="subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount",
        advanced_threat_protection_setting={
            "id": "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount/providers/Microsoft.Security/advancedThreatProtectionSettings/current",
            "name": "current",
            "properties": {"isEnabled": True},
            "type": "Microsoft.Security/advancedThreatProtectionSettings",
        },
    )
    print(response)


# x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-01-01/examples/AdvancedThreatProtection/PutAdvancedThreatProtectionSettings_example.json
if __name__ == "__main__":
    main()
