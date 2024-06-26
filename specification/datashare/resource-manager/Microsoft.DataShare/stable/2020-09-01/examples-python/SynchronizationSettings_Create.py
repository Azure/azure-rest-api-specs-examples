from azure.identity import DefaultAzureCredential
from azure.mgmt.datashare import DataShareManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-datashare
# USAGE
    python synchronization_settings_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataShareManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="433a8dfd-e5d5-4e77-ad86-90acdc75eb1a",
    )

    response = client.synchronization_settings.create(
        resource_group_name="SampleResourceGroup",
        account_name="Account1",
        share_name="Share1",
        synchronization_setting_name="Dataset1",
        synchronization_setting={
            "kind": "ScheduleBased",
            "properties": {"recurrenceInterval": "Day", "synchronizationTime": "2018-11-14T04:47:52.9614956Z"},
        },
    )
    print(response)


# x-ms-original-file: specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/SynchronizationSettings_Create.json
if __name__ == "__main__":
    main()
