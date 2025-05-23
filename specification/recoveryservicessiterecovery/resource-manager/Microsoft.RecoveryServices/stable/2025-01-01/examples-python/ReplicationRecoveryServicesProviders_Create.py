from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservicessiterecovery import SiteRecoveryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservicessiterecovery
# USAGE
    python replication_recovery_services_providers_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SiteRecoveryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="cb53d0c3-bd59-4721-89bc-06916a9147ef",
        resource_group_name="resourcegroup1",
        resource_name="migrationvault",
    )

    response = client.replication_recovery_services_providers.begin_create(
        fabric_name="vmwarefabric1",
        provider_name="vmwareprovider1",
        add_provider_input={
            "properties": {
                "authenticationIdentityInput": {
                    "aadAuthority": "https://login.microsoftonline.com",
                    "applicationId": "f66fce08-c0c6-47a1-beeb-0ede5ea94f90",
                    "audience": "https://microsoft.onmicrosoft.com/cf19e349-644c-4c6a-bcae-9c8f35357874",
                    "objectId": "141360b8-5686-4240-a027-5e24e6affeba",
                    "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
                },
                "machineName": "vmwareprovider1",
                "resourceAccessIdentityInput": {
                    "aadAuthority": "https://login.microsoftonline.com",
                    "applicationId": "f66fce08-c0c6-47a1-beeb-0ede5ea94f90",
                    "audience": "https://microsoft.onmicrosoft.com/cf19e349-644c-4c6a-bcae-9c8f35357874",
                    "objectId": "141360b8-5686-4240-a027-5e24e6affeba",
                    "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationRecoveryServicesProviders_Create.json
if __name__ == "__main__":
    main()
