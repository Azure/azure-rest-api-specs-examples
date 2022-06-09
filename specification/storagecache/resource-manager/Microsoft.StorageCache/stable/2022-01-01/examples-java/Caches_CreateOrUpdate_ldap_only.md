```java
import com.azure.resourcemanager.storagecache.models.CacheDirectorySettings;
import com.azure.resourcemanager.storagecache.models.CacheEncryptionSettings;
import com.azure.resourcemanager.storagecache.models.CacheSecuritySettings;
import com.azure.resourcemanager.storagecache.models.CacheSku;
import com.azure.resourcemanager.storagecache.models.CacheUsernameDownloadSettings;
import com.azure.resourcemanager.storagecache.models.CacheUsernameDownloadSettingsCredentials;
import com.azure.resourcemanager.storagecache.models.KeyVaultKeyReference;
import com.azure.resourcemanager.storagecache.models.KeyVaultKeyReferenceSourceVault;
import com.azure.resourcemanager.storagecache.models.NfsAccessPolicy;
import com.azure.resourcemanager.storagecache.models.NfsAccessRule;
import com.azure.resourcemanager.storagecache.models.NfsAccessRuleAccess;
import com.azure.resourcemanager.storagecache.models.NfsAccessRuleScope;
import com.azure.resourcemanager.storagecache.models.UsernameSource;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Caches CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_CreateOrUpdate_ldap_only.json
     */
    /**
     * Sample code: Caches_CreateOrUpdate_ldap_only.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesCreateOrUpdateLdapOnly(
        com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager
            .caches()
            .define("sc1")
            .withRegion("westus")
            .withExistingResourceGroup("scgroup")
            .withTags(mapOf("Dept", "Contoso"))
            .withSku(new CacheSku().withName("Standard_2G"))
            .withCacheSizeGB(3072)
            .withSubnet(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/sub1")
            .withEncryptionSettings(
                new CacheEncryptionSettings()
                    .withKeyEncryptionKey(
                        new KeyVaultKeyReference()
                            .withKeyUrl("https://keyvault-cmk.vault.azure.net/keys/key2048/test")
                            .withSourceVault(
                                new KeyVaultKeyReferenceSourceVault()
                                    .withId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"))))
            .withSecuritySettings(
                new CacheSecuritySettings()
                    .withAccessPolicies(
                        Arrays
                            .asList(
                                new NfsAccessPolicy()
                                    .withName("default")
                                    .withAccessRules(
                                        Arrays
                                            .asList(
                                                new NfsAccessRule()
                                                    .withScope(NfsAccessRuleScope.DEFAULT)
                                                    .withAccess(NfsAccessRuleAccess.RW)
                                                    .withSuid(false)
                                                    .withSubmountAccess(true)
                                                    .withRootSquash(false))))))
            .withDirectoryServicesSettings(
                new CacheDirectorySettings()
                    .withUsernameDownload(
                        new CacheUsernameDownloadSettings()
                            .withExtendedGroups(true)
                            .withUsernameSource(UsernameSource.LDAP)
                            .withLdapServer("192.0.2.12")
                            .withLdapBaseDN("dc=contosoad,dc=contoso,dc=local")
                            .withCredentials(
                                new CacheUsernameDownloadSettingsCredentials()
                                    .withBindDn("cn=ldapadmin,dc=contosoad,dc=contoso,dc=local")
                                    .withBindPassword("<bindPassword>"))))
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagecache_1.0.0-beta.5/sdk/storagecache/azure-resourcemanager-storagecache/README.md) on how to add the SDK to your project and authenticate.
