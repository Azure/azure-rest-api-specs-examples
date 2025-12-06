
import com.azure.resourcemanager.netapp.models.CacheProperties;
import com.azure.resourcemanager.netapp.models.EncryptionKeySource;
import com.azure.resourcemanager.netapp.models.LdapServerType;
import com.azure.resourcemanager.netapp.models.LdapState;
import com.azure.resourcemanager.netapp.models.OriginClusterInformation;
import java.util.Arrays;

/**
 * Samples for Caches CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Caches_CreateOrUpdate.json
     */
    /**
     * Sample code: Caches_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().define("cache1").withRegion("eastus").withExistingCapacityPool("myRG", "account1", "pool1")
            .withProperties(new CacheProperties().withFilepath("cache-west-us2-01").withSize(107374182400L)
                .withCacheSubnetResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/cacheVnet/subnets/cacheSubnet1")
                .withPeeringSubnetResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/icLifVnet/subnets/peeringSubnet1")
                .withEncryptionKeySource(EncryptionKeySource.MICROSOFT_NET_APP).withLdap(LdapState.ENABLED)
                .withLdapServerType(LdapServerType.OPEN_LDAP)
                .withOriginClusterInformation(new OriginClusterInformation().withPeerClusterName("cluster1")
                    .withPeerAddresses(Arrays.asList("192.0.2.10", "192.0.2.11")).withPeerVserverName("vserver1")
                    .withPeerVolumeName("originvol1")))
            .create();
    }
}
