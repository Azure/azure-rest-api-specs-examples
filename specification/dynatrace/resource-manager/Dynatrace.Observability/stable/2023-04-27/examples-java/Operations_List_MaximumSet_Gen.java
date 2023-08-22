/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void operationsListMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
