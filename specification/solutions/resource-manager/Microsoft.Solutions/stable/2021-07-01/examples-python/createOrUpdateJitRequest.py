from azure.identity import DefaultAzureCredential
from azure.mgmt.managedapplications import ManagedApplicationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managedapplications
# USAGE
    python create_or_update_jit_request.py

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

    response = client.jit_requests.begin_create_or_update(
        resource_group_name="rg",
        jit_request_name="myJitRequest",
        parameters={
            "properties": {
                "applicationResourceId": "/subscriptions/00c76877-e316-48a7-af60-4a09fec9d43f/resourceGroups/52F30DB2/providers/Microsoft.Solutions/applications/7E193158",
                "jitAuthorizationPolicies": [
                    {
                        "principalId": "1db8e132e2934dbcb8e1178a61319491",
                        "roleDefinitionId": "ecd05a23-931a-4c38-a52b-ac7c4c583334",
                    }
                ],
                "jitSchedulingPolicy": {
                    "duration": "PT8H",
                    "startTime": "2021-04-22T05:48:30.6661804Z",
                    "type": "Once",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateJitRequest.json
if __name__ == "__main__":
    main()
