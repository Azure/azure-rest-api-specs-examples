
/**
 * Samples for KubernetesVersions ListBySubscriptionLocationResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/KubernetesVersions_ListBySubscriptionLocationResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: KubernetesVersions_ListBySubscriptionLocationResource_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void kubernetesVersionsListBySubscriptionLocationResourceMaximumSet(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.kubernetesVersions().listBySubscriptionLocationResource("westus2", com.azure.core.util.Context.NONE);
    }
}
