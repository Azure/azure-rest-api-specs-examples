from azure.identity import DefaultAzureCredential
from azure.mgmt.astro import AstroMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-astro
# USAGE
    python organizations_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AstroMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="43454B17-172A-40FE-80FA-549EA23D12B3",
    )

    response = client.organizations.begin_update(
        resource_group_name="rgastronomer",
        organization_name="6.",
        properties={
            "identity": {"type": "None", "userAssignedIdentities": {}},
            "properties": {
                "partnerOrganizationProperties": {
                    "organizationId": "lrtmbkvyvvoszhjevohkmyjhfyty",
                    "organizationName": "U2P_",
                    "singleSignOnProperties": {
                        "aadDomains": ["kfbleh"],
                        "enterpriseAppId": "mklfypyujwumgwdzae",
                        "provisioningState": "Succeeded",
                        "singleSignOnState": "Initial",
                        "singleSignOnUrl": "ymmtzkyghvinvhgnqlzwrr",
                    },
                    "workspaceId": "xsepuskdhejaadusyxq",
                    "workspaceName": "L.-y_--:",
                },
                "user": {
                    "emailAddress": ".K_@e7N-g1.xjqnbPs",
                    "firstName": "qeuofehzypzljgcuysugefbgxde",
                    "lastName": "g",
                    "phoneNumber": "aqpyxznvqpgkzohevynofrjdfgoo",
                    "upn": "uwtprzdfpsqmktx",
                },
            },
            "tags": {"key1474": "bqqyipxnbbxryhznyaosmtpo"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Organizations_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
