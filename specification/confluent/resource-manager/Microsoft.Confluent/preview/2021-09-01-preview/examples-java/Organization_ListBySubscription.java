import com.azure.core.util.Context;

/** Samples for Organization List. */
public final class Main {
    /*
     * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_ListBySubscription.json
     */
    /**
     * Sample code: Organization_ListBySubscription.
     *
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationListBySubscription(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().list(Context.NONE);
    }
}
