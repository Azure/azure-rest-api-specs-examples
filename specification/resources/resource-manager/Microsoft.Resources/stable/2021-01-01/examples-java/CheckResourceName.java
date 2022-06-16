import com.azure.core.util.Context;

/** Samples for ResourceName CheckResourceName. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/CheckResourceName.json
     */
    /**
     * Sample code: CheckValidityForAResourceName.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkValidityForAResourceName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .subscriptionClient()
            .getResourceNames()
            .checkResourceNameWithResponse(null, Context.NONE);
    }
}
