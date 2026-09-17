
import com.azure.resourcemanager.network.fluent.models.IpamPoolInner;
import com.azure.resourcemanager.network.models.IpamPoolProperties;
import java.util.Arrays;

/**
 * Samples for IpamPools Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/IpamPools_CreateWithAllocationBounds.json
     */
    /**
     * Sample code: Create/Update the Pool resource with allocation size bounds.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createUpdateThePoolResourceWithAllocationSizeBounds(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().create("rg1", "TestNetworkManager", "TestPool",
            new IpamPoolInner().withLocation("eastus")
                .withProperties(new IpamPoolProperties().withDescription("Test description.").withParentPoolName("")
                    .withAddressPrefixes(Arrays.asList("10.0.0.0/24")).withMinAllocationSize("16")
                    .withMaxAllocationSize("256")),
            null, com.azure.core.util.Context.NONE);
    }
}
