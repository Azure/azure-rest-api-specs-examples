from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.qumulo import QumuloMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-qumulo
# USAGE
    python file_systems_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = QumuloMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="382E8C7A-AC80-4D70-8580-EFE99537B9B7",
    )

    response = client.file_systems.begin_create_or_update(
        resource_group_name="rgQumulo",
        file_system_name="hfcmtgaes",
        resource={
            "identity": {"type": "None", "userAssignedIdentities": {"key7679": {}}},
            "location": "pnb",
            "properties": {
                "adminPassword": "fakeTestSecretPlaceholder",
                "availabilityZone": "eqdvbdiuwmhhzqzmksmwllpddqquwt",
                "clusterLoginUrl": "ykaynsjvhihdthkkvvodjrgc",
                "delegatedSubnetId": "jykmxrf",
                "marketplaceDetails": {
                    "marketplaceSubscriptionId": "xaqtkloiyovmexqhn",
                    "marketplaceSubscriptionStatus": "PendingFulfillmentStart",
                    "offerId": "s",
                    "planId": "fwtpz",
                    "publisherId": "czxcfrwodazyaft",
                    "termUnit": "cfwwczmygsimcyvoclcw",
                },
                "privateIPs": ["gzken"],
                "storageSku": "yhyzby",
                "userDetails": {"email": "aqsnzyroo"},
            },
            "tags": {"key7090": "rurrdiaqp"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
