class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        boxTypes.sort(key=lambda x : x[1], reverse = True)
        result = 0
        truck = 0
        print(boxTypes)
        
        for box in boxTypes:
            truck += box[0]
            if truck >= truckSize:
                result += (truckSize - (truck - box[0])) * box[1]
                break
            else:    
                result += box[0] * box[1]
            
        return result
    
    
