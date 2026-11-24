
import com.azure.resourcemanager.servicebus.models.FailOver;

/**
 * Samples for Namespaces Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNamespaceFailover.json
     */
    /**
     * Sample code: NameSpaceCreate.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceCreate(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().failover("ResurceGroupSample", "NamespaceGeoDRFailoverSample",
            new FailOver().withPrimaryLocation("centralus").withForce(true), com.azure.core.util.Context.NONE);
    }
}
