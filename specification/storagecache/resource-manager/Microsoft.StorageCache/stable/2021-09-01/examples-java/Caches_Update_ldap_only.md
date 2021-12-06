Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagecache_1.0.0-beta.4/sdk/storagecache/azure-resourcemanager-storagecache/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storagecache.models.Cache;
import com.azure.resourcemanager.storagecache.models.CacheDirectorySettings;
import com.azure.resourcemanager.storagecache.models.CacheNetworkSettings;
import com.azure.resourcemanager.storagecache.models.CacheSecuritySettings;
import com.azure.resourcemanager.storagecache.models.CacheUsernameDownloadSettings;
import com.azure.resourcemanager.storagecache.models.CacheUsernameDownloadSettingsCredentials;
import com.azure.resourcemanager.storagecache.models.NfsAccessPolicy;
import com.azure.resourcemanager.storagecache.models.NfsAccessRule;
import com.azure.resourcemanager.storagecache.models.NfsAccessRuleAccess;
import com.azure.resourcemanager.storagecache.models.NfsAccessRuleScope;
import com.azure.resourcemanager.storagecache.models.UsernameSource;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Caches Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/Caches_Update_ldap_only.json
     */
    /**
     * Sample code: Caches_Update_ldap_only.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesUpdateLdapOnly(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        Cache resource = manager.caches().getByResourceGroupWithResponse("scgroup", "sc1", Context.NONE).getValue();
        resource
            .update()
            .withTags(mapOf("Dept", "Contoso"))
            .withNetworkSettings(
                new CacheNetworkSettings()
                    .withMtu(1500)
                    .withDnsServers(Arrays.asList("10.1.22.33", "10.1.12.33"))
                    .withDnsSearchDomain("contoso.com")
                    .withNtpServer("time.contoso.com"))
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
                                                    .withRootSquash(false))),
                                new NfsAccessPolicy()
                                    .withName("restrictive")
                                    .withAccessRules(
                                        Arrays
                                            .asList(
                                                new NfsAccessRule()
                                                    .withScope(NfsAccessRuleScope.HOST)
                                                    .withFilter("10.99.3.145")
                                                    .withAccess(NfsAccessRuleAccess.RW)
                                                    .withSuid(true)
                                                    .withSubmountAccess(true)
                                                    .withRootSquash(false),
                                                new NfsAccessRule()
                                                    .withScope(NfsAccessRuleScope.NETWORK)
                                                    .withFilter("10.99.1.0/24")
                                                    .withAccess(NfsAccessRuleAccess.RW)
                                                    .withSuid(true)
                                                    .withSubmountAccess(true)
                                                    .withRootSquash(false),
                                                new NfsAccessRule()
                                                    .withScope(NfsAccessRuleScope.DEFAULT)
                                                    .withAccess(NfsAccessRuleAccess.NO)
                                                    .withSuid(false)
                                                    .withSubmountAccess(true)
                                                    .withRootSquash(true)
                                                    .withAnonymousUid("65534")
                                                    .withAnonymousGid("65534"))))))
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
            .apply();
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
