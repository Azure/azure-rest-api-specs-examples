
import com.azure.resourcemanager.netapp.models.ExternalCacheConnectionProperties;

/**
 * Samples for ExternalCacheConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalCacheConnections_CreateOrUpdate.json
     */
    /**
     * Sample code: ExternalCacheConnections_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        externalCacheConnectionsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalCacheConnections().define("contosoCacheConnection").withRegion("eastus")
            .withExistingResourceGroup("contosoNetAppRg")
            .withProperties(new ExternalCacheConnectionProperties().withOriginVolumeResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/netAppAccounts/contosoAccount/capacityPools/contosoPool/volumes/contosoVolume")
                .withExternalCacheName("contosoCache")
                .withExternalSvmResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.NetApp/externalClusters/contosoExternalCluster/externalSvms/contosoExternalSvm")
                .withClusterPeerSubnetResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoNetAppRg/providers/Microsoft.Network/virtualNetworks/contosoVnet/subnets/contosoSubnet"))
            .create();
    }
}
