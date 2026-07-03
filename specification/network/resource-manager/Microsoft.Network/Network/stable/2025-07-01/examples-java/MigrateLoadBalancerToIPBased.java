
import com.azure.resourcemanager.network.models.MigrateLoadBalancerToIpBasedRequest;
import java.util.Arrays;

/**
 * Samples for LoadBalancers MigrateToIpBased.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/MigrateLoadBalancerToIPBased.json
     */
    /**
     * Sample code: Migrate load balancer to IP Based.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void migrateLoadBalancerToIPBased(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancers().migrateToIpBasedWithResponse("rg1", "lb1",
            new MigrateLoadBalancerToIpBasedRequest().withPools(Arrays.asList("pool1", "pool2")),
            com.azure.core.util.Context.NONE);
    }
}
