
import com.azure.resourcemanager.eventhubs.models.FailOver;

/**
 * Samples for Namespaces Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNamespaceFailover.json
     */
    /**
     * Sample code: NameSpaceCreate.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void nameSpaceCreate(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().failover("ResurceGroupSample", "NamespaceGeoDRFailoverSample",
            new FailOver().withPrimaryLocation("centralus").withForce(true), com.azure.core.util.Context.NONE);
    }
}
