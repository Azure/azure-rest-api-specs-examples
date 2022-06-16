import com.azure.core.util.Context;

/** Samples for Organization ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_ListByResourceGroup.json
     */
    /**
     * Sample code: Organization_ListByResourceGroup.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationListByResourceGroup(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
