from azure.identity import DefaultAzureCredential

from azure.mgmt.lambdatesthyperexecute import LambdaTestHyperExecuteMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-lambdatesthyperexecute
# USAGE
    python organizations_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LambdaTestHyperExecuteMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.organizations.begin_create_or_update(
        resource_group_name="rgopenapi",
        organizationname="testorg",
        resource={
            "identity": {"type": "None", "userAssignedIdentities": {}},
            "location": "cvymsrlt",
            "properties": {
                "marketplace": {
                    "offerDetails": {
                        "offerId": "fmljkvoivqmfdiwsu",
                        "planId": "ssjlabxexw",
                        "planName": "mrguqu",
                        "publisherId": "ufwwpzit",
                        "termId": "hxkszxfscsyefeuunyyfskhibr",
                        "termUnit": "acvhavsffebfivyaxhxxsaqzt",
                    },
                    "subscriptionId": "zetdxwryjgcsnosezfpovkpvgvim",
                    "subscriptionStatus": "PendingFulfillmentStart",
                },
                "partnerProperties": {"licensesSubscribed": 7},
                "singleSignOnProperties": {
                    "aadDomains": ["hrgguokssgyrfdzliyjmovtelfu"],
                    "enterpriseAppId": "sonpowym",
                    "state": "Initial",
                    "type": "Saml",
                    "url": "qlshnxrfcdpjcpkxxisrn",
                },
                "user": {
                    "emailAddress": "joe@example.com",
                    "firstName": "ssnzyujsrszbptndzeoqzrmbufrhgp",
                    "lastName": "nsfylyvdyrtfzfeehmji",
                    "phoneNumber": "jkevskjaaylbwjzofkzmxdysejsoir",
                    "upn": "tfqolz",
                },
            },
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-02-01/Organizations_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
