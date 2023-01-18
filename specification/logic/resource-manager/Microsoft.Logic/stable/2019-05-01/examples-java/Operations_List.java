/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Operations_List.json
     */
    /**
     * Sample code: List Microsoft.Logic operations.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listMicrosoftLogicOperations(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
