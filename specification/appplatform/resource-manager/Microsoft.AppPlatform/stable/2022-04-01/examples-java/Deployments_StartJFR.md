```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.models.DiagnosticParameters;

/** Samples for Deployments StartJfr. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_StartJFR.json
     */
    /**
     * Sample code: Deployments_StartJFR.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsStartJFR(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .startJfr(
                "myResourceGroup",
                "myservice",
                "myapp",
                "mydeployment",
                new DiagnosticParameters()
                    .withAppInstance("myappinstance")
                    .withFilePath("/byos/diagnose")
                    .withDuration("60s"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
