from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python exa_infra_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OracleDatabaseMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.cloud_exadata_infrastructures.begin_create_or_update(
        resource_group_name="rg000",
        cloudexadatainfrastructurename="infra1",
        resource={
            "location": "eastus",
            "properties": {"computeCount": 100, "displayName": "infra 1", "shape": "EXADATA.X9M", "storageCount": 10},
            "tags": {"tagK1": "tagV1"},
            "zones": ["1"],
        },
    ).result()
    print(response)


# x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_create.json
if __name__ == "__main__":
    main()
