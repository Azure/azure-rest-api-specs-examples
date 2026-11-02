
import com.azure.resourcemanager.netapp.models.CacheProperties;
import com.azure.resourcemanager.netapp.models.EncryptionKeySource;
import com.azure.resourcemanager.netapp.models.OriginClusterInformation;
import com.azure.resourcemanager.netapp.models.SmbAccessBasedEnumeration;
import com.azure.resourcemanager.netapp.models.SmbEncryptionState;
import com.azure.resourcemanager.netapp.models.SmbNonBrowsable;
import com.azure.resourcemanager.netapp.models.SmbOplocks;
import com.azure.resourcemanager.netapp.models.SmbSettings;
import java.util.Arrays;

/**
 * Samples for Caches CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/Caches_CreateOrUpdate_SMB.json
     */
    /**
     * Sample code: Caches_CreateOrUpdate_SMB.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesCreateOrUpdateSMB(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().define("cache1").withRegion("eastus").withExistingCapacityPool("myRG", "account1", "pool1")
            .withProperties(new CacheProperties().withFilePath("cache-smb-01").withSize(107374182400L)
                .withCacheSubnetResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/cacheVnet/subnets/cacheSubnet1")
                .withPeeringSubnetResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/icLifVnet/subnets/peeringSubnet1")
                .withSmbSettings(new SmbSettings().withSmbEncryption(SmbEncryptionState.ENABLED)
                    .withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.ENABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.ENABLED).withSmbOplocks(SmbOplocks.ENABLED))
                .withEncryptionKeySource(EncryptionKeySource.MICROSOFT_NET_APP)
                .withOriginClusterInformation(new OriginClusterInformation().withPeerClusterName("cluster1")
                    .withPeerAddresses(Arrays.asList("192.0.2.10", "192.0.2.11")).withPeerVserverName("vserver1")
                    .withPeerVolumeName("originvol1")))
            .create();
    }
}
