import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.models.DiagnosticParameters;

/** Samples for Deployments GenerateHeapDump. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Deployments_GenerateHeapDump.json
     */
    /**
     * Sample code: Deployments_GenerateHeapDump.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsGenerateHeapDump(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .generateHeapDump(
                "myResourceGroup",
                "myservice",
                "myapp",
                "mydeployment",
                new DiagnosticParameters().withAppInstance("myappinstance").withFilePath("/byos/diagnose"),
                Context.NONE);
    }
}
