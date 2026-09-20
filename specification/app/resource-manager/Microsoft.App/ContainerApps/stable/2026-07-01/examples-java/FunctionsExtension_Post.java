
/**
 * Samples for FunctionsExtension InvokeFunctionsHost.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/FunctionsExtension_Post.json
     */
    /**
     * Sample code: Invoke Functions host using Functions Extension API.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void invokeFunctionsHostUsingFunctionsExtensionAPI(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.functionsExtensions().invokeFunctionsHostWithResponse("rg", "testcontainerApp0",
            "testcontainerApp0-pjxhsye", "testcontainerApp0", com.azure.core.util.Context.NONE);
    }
}
