from azure.identity import DefaultAzureCredential

from azure.mgmt.secretsstoreextension import SecretsStoreExtensionMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-secretsstoreextension
# USAGE
    python secret_syncs_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecretsStoreExtensionMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.secret_syncs.begin_update(
        resource_group_name="rg-ssc-example",
        secret_sync_name="secretsync-ssc-example",
        properties={
            "properties": {
                "forceSynchronization": "arbitrarystring",
                "kubernetesSecretType": "Opaque",
                "objectSecretMapping": [
                    {
                        "sourcePath": "ssrzmbvdiomkvzrdsyilwlfzicfydnbjwjsnohrppkukjddrunfslkrnexunuckmghixdssposvndpiqchpqrkjuqbapoisvqdvgstvdonsmlpsmticfvuhqlofpaxfdg",
                        "targetKey": "lojegeqiqfjxyblfxhxloccqzwgpgcelrwqwsbsltcjvjvceejgdrmxhenokxrylhynkltvqntjcvujjrppzvcxyivxeksgmzhifrcklizbpntdepzdl",
                    }
                ],
                "secretProviderClassName": "jttlpenhtpxfrrlxdsmqqvmvtmgqrficvqngkggjwciilrexenlstxncyvkqcydxrivkioujssncoaiysdklfouukczzdbxniipbyiqsarqaespuqrbbydwtdaulllostoomntkadklihemfpeffvuyvyilequiqewzspaootvkibrynbqrsbiptjdhywynvydaadprdc",
                "serviceAccountName": "fcldqfdfpktndlntuoxicsftelhefevovmlycflfwzckvamiqjnjugandqaqqeccsbzztfmmeunvhsafgerbcsdbnmsyqivygornebbkusuvphwghgouxvcbvmbydqjzoxextnyowsnyymadniwdrrxtogeveldpejixmsrzzfqkquaxdpzwvecevqwasxgxxchrfa",
            },
            "tags": {"example-tag": "example-tag-value"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-08-21-preview/SecretSyncs_Update_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
