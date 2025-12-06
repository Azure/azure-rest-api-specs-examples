
import com.azure.resourcemanager.netapp.models.ElasticCapacityPoolProperties;
import com.azure.resourcemanager.netapp.models.ElasticServiceLevel;
import java.util.Arrays;

/**
 * Samples for ElasticCapacityPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticCapacityPools_CreateOrUpdate.json
     */
    /**
     * Sample code: ElasticCapacityPools_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticCapacityPoolsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticCapacityPools().define("pool1").withRegion("eastus")
            .withExistingElasticAccount("myRG", "account1")
            .withProperties(new ElasticCapacityPoolProperties().withSize(4398046511104L)
                .withServiceLevel(ElasticServiceLevel.ZONE_REDUNDANT)
                .withSubnetResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                .withActiveDirectoryConfigResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/activeDirectoryConfigs/activeDirectoryConfig1"))
            .withZones(Arrays.asList("1", "2", "3")).create();
    }
}
