from azure.identity import DefaultAzureCredential

from azure.mgmt.neonpostgres import NeonPostgresMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-neonpostgres
# USAGE
    python organizations_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NeonPostgresMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.organizations.begin_create_or_update(
        resource_group_name="rgneon",
        organization_name="XB-.:",
        resource={
            "location": "upxxgikyqrbnv",
            "properties": {
                "companyDetails": {
                    "businessPhone": "hbeb",
                    "companyName": "uxn",
                    "country": "lpajqzptqchuko",
                    "domain": "krjldeakhwiepvs",
                    "numberOfEmployees": 23,
                    "officeAddress": "chpkrlpmfslmawgunjxdllzcrctykq",
                },
                "marketplaceDetails": {
                    "offerDetails": {
                        "offerId": "bunyeeupoedueofwrzej",
                        "planId": "nlbfiwtslenfwek",
                        "planName": "ljbmgpkfqklaufacbpml",
                        "publisherId": "hporaxnopmolttlnkbarw",
                        "termId": "aedlchikwqckuploswthvshe",
                        "termUnit": "qbcq",
                    },
                    "subscriptionId": "yxmkfivp",
                    "subscriptionStatus": "PendingFulfillmentStart",
                },
                "partnerOrganizationProperties": {
                    "organizationId": "nrhvoqzulowcunhmvwfgjcaibvwcl",
                    "organizationName": "2__.-",
                    "singleSignOnProperties": {
                        "aadDomains": ["kndszgrwzbvvlssvkej"],
                        "enterpriseAppId": "fpibacregjfncfdsojs",
                        "singleSignOnState": "Initial",
                        "singleSignOnUrl": "tmojh",
                    },
                },
                "userDetails": {
                    "emailAddress": "3i_%@w8-y.H-p.tvj.dG",
                    "firstName": "buwwe",
                    "lastName": "escynjpynkoox",
                    "phoneNumber": "dlrqoowumy",
                    "upn": "fwedjamgwwrotcjaucuzdwycfjdqn",
                },
            },
            "tags": {"key2099": "omjjymaqtrqzksxszhzgyl"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-08-01-preview/Organizations_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
