from azure.identity import DefaultAzureCredential

from azure.mgmt.connectedcache import ConnectedCacheMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-connectedcache
# USAGE
    python enterprise_customer_operations_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ConnectedCacheMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.enterprise_customer_operations.begin_create_or_update(
        resource_group_name="rgConnectedCache",
        customer_resource_name="l",
        resource={
            "location": "zdzhhkjyogrqxwihkifnmeyhwpujbr",
            "properties": {
                "error": {},
                "statusCode": "oldkroffqtkryqffpsi",
                "statusDetails": "lhwvcz",
                "statusText": "bs",
            },
            "tags": {"key4215": "zjbszvlzf"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2023-05-01-preview/EnterpriseCustomerOperations_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
