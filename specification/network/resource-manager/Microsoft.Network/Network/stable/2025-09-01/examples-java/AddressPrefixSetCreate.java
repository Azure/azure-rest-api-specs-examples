
import com.azure.resourcemanager.network.fluent.models.AddressPrefixSetInner;
import com.azure.resourcemanager.network.models.AddressPrefixSetPropertiesFormat;
import java.util.Arrays;

/**
 * Samples for AddressPrefixSets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/AddressPrefixSetCreate.json
     */
    /**
     * Sample code: Create address prefix set.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createAddressPrefixSet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAddressPrefixSets().createOrUpdate("rg1", "test-asg", "test-prefix-set",
            new AddressPrefixSetInner().withProperties(new AddressPrefixSetPropertiesFormat()
                .withAddressPrefixes(Arrays.asList("10.0.0.0/16", "192.168.1.0/24", "2001:db8::/32"))),
            com.azure.core.util.Context.NONE);
    }
}
