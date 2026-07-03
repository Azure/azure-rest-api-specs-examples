
import com.azure.resourcemanager.network.fluent.models.IpamPoolInner;
import com.azure.resourcemanager.network.models.IpamPoolProperties;
import java.util.Arrays;

/**
 * Samples for IpamPools Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/IpamPools_Create.json
     */
    /**
     * Sample code: IpamPools_Create.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void ipamPoolsCreate(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getIpamPools().create("rg1", "TestNetworkManager", "TestPool",
            new IpamPoolInner().withLocation("eastus")
                .withProperties(new IpamPoolProperties().withDescription("Test description.").withParentPoolName("")
                    .withAddressPrefixes(Arrays.asList("10.0.0.0/24"))),
            null, com.azure.core.util.Context.NONE);
    }
}
