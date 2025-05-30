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
        organization_name="test-org",
        resource={
            "location": "kcdph",
            "properties": {
                "companyDetails": {
                    "businessPhone": "hucxvzcvpaupqjkgb",
                    "companyName": "xtul",
                    "country": "ycmyjdcpyjieemfrthfyxdlvn",
                    "domain": "snoshqumfsthyofpnrsgyjhszvgtj",
                    "numberOfEmployees": 12,
                    "officeAddress": "icirtoqmmozijk",
                },
                "marketplaceDetails": {
                    "offerDetails": {
                        "offerId": "qscggwfdnippiwrrnmuscg",
                        "planId": "sveqoxtdwxutxmtniuufyrdu",
                        "planName": "t",
                        "publisherId": "eibghzuyqsksouwlgqphhmuxeqeigf",
                        "termId": "uptombvymytfonj",
                        "termUnit": "jnxhyql",
                    },
                    "subscriptionId": "xfahbbbzwlcwhhjbxarnwfcy",
                    "subscriptionStatus": "PendingFulfillmentStart",
                },
                "partnerOrganizationProperties": {
                    "organizationId": "hzejhmftwsruhwspvtwoy",
                    "organizationName": "entity-name",
                    "singleSignOnProperties": {
                        "aadDomains": ["mdzbelaiphukhe"],
                        "enterpriseAppId": "urtjzjfr",
                        "singleSignOnState": "Initial",
                        "singleSignOnUrl": "gcmlwvtxcsjozitm",
                    },
                },
                "projectProperties": {
                    "attributes": [{"name": "trhvzyvaqy", "value": "evpkgsskyavybxwwssm"}],
                    "branch": {
                        "attributes": [{"name": "trhvzyvaqy", "value": "evpkgsskyavybxwwssm"}],
                        "databaseName": "duhxebzhd",
                        "databases": [
                            {
                                "attributes": [{"name": "trhvzyvaqy", "value": "evpkgsskyavybxwwssm"}],
                                "branchId": "orfdwdmzvfvlnrgussvcvoek",
                                "entityName": "entity-name",
                                "ownerName": "odmbeg",
                            }
                        ],
                        "endpoints": [
                            {
                                "attributes": [{"name": "trhvzyvaqy", "value": "evpkgsskyavybxwwssm"}],
                                "branchId": "rzsyrhpfbydxtfkpaa",
                                "endpointType": "read_only",
                                "entityName": "entity-name",
                                "projectId": "rtvdeeflqzlrpfzhjqhcsfbldw",
                            }
                        ],
                        "entityName": "entity-name",
                        "parentId": "entity-id",
                        "projectId": "oik",
                        "roleName": "qrrairsupyosxnqotdwhbpc",
                        "roles": [
                            {
                                "attributes": [{"name": "trhvzyvaqy", "value": "evpkgsskyavybxwwssm"}],
                                "branchId": "wxbojkmdgaggkfiwqfakdkbyztm",
                                "entityName": "entity-name",
                                "isSuperUser": True,
                                "permissions": ["myucqecpjriewzohxvadgkhiudnyx"],
                            }
                        ],
                    },
                    "databases": [
                        {
                            "attributes": [{"name": "trhvzyvaqy", "value": "evpkgsskyavybxwwssm"}],
                            "branchId": "orfdwdmzvfvlnrgussvcvoek",
                            "entityName": "entity-name",
                            "ownerName": "odmbeg",
                        }
                    ],
                    "defaultEndpointSettings": {"autoscalingLimitMaxCu": 20, "autoscalingLimitMinCu": 26},
                    "endpoints": [
                        {
                            "attributes": [{"name": "trhvzyvaqy", "value": "evpkgsskyavybxwwssm"}],
                            "branchId": "rzsyrhpfbydxtfkpaa",
                            "endpointType": "read_only",
                            "entityName": "entity-name",
                            "projectId": "rtvdeeflqzlrpfzhjqhcsfbldw",
                        }
                    ],
                    "entityName": "entity-name",
                    "historyRetention": 7,
                    "pgVersion": 10,
                    "regionId": "tlcltldfrnxh",
                    "roles": [
                        {
                            "attributes": [{"name": "trhvzyvaqy", "value": "evpkgsskyavybxwwssm"}],
                            "branchId": "wxbojkmdgaggkfiwqfakdkbyztm",
                            "entityName": "entity-name",
                            "isSuperUser": True,
                            "permissions": ["myucqecpjriewzohxvadgkhiudnyx"],
                        }
                    ],
                    "storage": 7,
                },
                "userDetails": {
                    "emailAddress": "test@contoso.com",
                    "firstName": "zhelh",
                    "lastName": "zbdhouyeozylnerrc",
                    "phoneNumber": "zmejenytglrmjnt",
                    "upn": "mixcikvxlnhkfugetqlngz",
                },
            },
            "tags": {"key8832": "rvukepuxkykdtqjtwk"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-03-01/Organizations_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
