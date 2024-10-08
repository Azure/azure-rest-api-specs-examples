from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.informaticadatamanagement import InformaticaDataMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-informaticadatamanagement
# USAGE
    python organizations_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = InformaticaDataMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="3599DA28-E346-4D9F-811E-189C0445F0FE",
    )

    response = client.organizations.begin_create_or_update(
        resource_group_name="rgopenapi",
        organization_name="C",
        resource={
            "location": "pamjoudtssthlbhrnfjidr",
            "properties": {
                "companyDetails": {
                    "business": "pucosrtjv",
                    "companyName": "xszcggknokhw",
                    "country": "gwkcpnwyaqc",
                    "domain": "utcxetzzpmbvwmjrvphqngvp",
                    "numberOfEmployees": 25,
                    "officeAddress": "sbttzwyajgdbsvipuiclbzvkcvwyil",
                },
                "informaticaProperties": {
                    "informaticaRegion": "zfqodqpbeflhedypiijdkc",
                    "organizationId": "wtdmhlwhkvgqdumaehgfgiqcxgnqpx",
                    "organizationName": "nomzbvwe",
                    "singleSignOnUrl": "https://contoso.com/singlesignon",
                },
                "linkOrganization": {"token": "jjfouhoqpumjvrdsfbimgcy"},
                "marketplaceDetails": {
                    "marketplaceSubscriptionId": "ovenlecocg",
                    "offerDetails": {
                        "offerId": "cwswcfwmzhjcoksmueukegwaptvpcmbfyvixfhvgwnjyblqivqdkkwkunkgimiopwwkvgnwclmajhuty",
                        "planId": "jfnemevyivtlxhectiutdavdgfyidolivuojumdzckp",
                        "planName": "iaoxgaitteuoqgujkgxbdgryaobtkjjecuvchwutntrvmuorikrbqqegmelenbewhakiysprrnovjixyxrikscaptrbapbdspu",
                        "publisherId": "zajxpfacudwongxjvnnuhhpygmnydchgowjccyuzsjonegmqxcqqpnzafanggowfqdixnnutyfvmvwrkx",
                        "termId": "tcvvsxdjnjlfmjhmvwklptdmxetnzydxyuhfqchoubmtoeqbchnfxoxqzezlgpxdnzyvzgkynjxzzgetkqccxvpzahxattluqdipvbdktqmndfefitzuifqjpschzlbvixnvznkmmgjwvkplfhemnapsewgqxggdzdokryhv",
                        "termUnit": "gjwmgevrblbosuogsvfspsgspetbnxaygkbelvadpgwiywl",
                    },
                },
                "provisioningState": "Accepted",
                "userDetails": {
                    "emailAddress": "7_-46@13D--3.m-4x-.11.c-9-.DHLYFc",
                    "firstName": "appvdclawzfjntdfdftjevlhvzropnxqtnypid",
                    "lastName": "nzirbvzmkxtbrlamyatlcszebxgcyncxoascojsmacwvjsjvn",
                    "phoneNumber": "fvcjylxlmhdnshsgywnzlyvshu",
                    "upn": "undljch",
                },
            },
            "tags": {"key8430": "cagshqtjlxtqqhdwtchokvxszybp"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
