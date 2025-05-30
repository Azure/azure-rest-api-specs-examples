from azure.identity import DefaultAzureCredential

from azure.mgmt.containerservicefleet import ContainerServiceFleetMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerservicefleet
# USAGE
    python auto_upgrade_profiles_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerServiceFleetMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.auto_upgrade_profiles.begin_create_or_update(
        resource_group_name="rg1",
        fleet_name="fleet1",
        auto_upgrade_profile_name="autoupgradeprofile1",
        resource={"properties": {"channel": "Stable"}},
    ).result()
    print(response)


# x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2025-03-01/examples/AutoUpgradeProfiles_CreateOrUpdate.json
if __name__ == "__main__":
    main()
