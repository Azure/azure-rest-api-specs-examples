from azure.identity import DefaultAzureCredential
from azure.mgmt.managedapplications import ManagedApplicationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managedapplications
# USAGE
    python get_jit_request.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedApplicationsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.jit_requests.get(
        resource_group_name="rg",
        jit_request_name="myJitRequest",
    )
    print(response)


# x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/getJitRequest.json
if __name__ == "__main__":
    main()
