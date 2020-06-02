﻿using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace Utils
{
	public static class Spatial
	{
		/// <summary>
		/// Return position above ground relatively from the prefab size
		/// Global position
		/// </summary>
		/// <param name="position"></param>
		/// <param name="prefabHeight">Prefab height needed in order to place well on top of ground</param>
		/// <param name="transform">Transform parent</param>
		/// <param name="layerMask">Layers to ignore</param>
		/// <returns></returns>
		public static Vector3 PositionAboveGround(this Vector3 position,
			float prefabHeight = 1f,
			Transform transform = null,
			LayerMask layerMask = default)
		{
			var p = position;
			if (transform != null) p += transform.position;


			// Current position is below ground
			if (Physics.Raycast(p, Vector3.up, out var hit, Mathf.Infinity, ~layerMask))
			{
				p.y += hit.distance + prefabHeight * 0.5f;
				return p;
			}


			// Current position is above ground
			if (Physics.Raycast(p, Vector3.down, out hit, Mathf.Infinity, ~layerMask))
			{
				p.y -= hit.distance - prefabHeight * 0.5f;
				return p;
			}

			// There is no ground above or below, outside map
			return Vector3.zero;
		}

		/// <summary>
		/// This function will find a position to spawn above ground and far enough from other objects of the given layer
		/// Returns Vector3 zero in case it couldn't find a free position
		/// </summary>
		/// <returns></returns>
		public static Vector3 RandomPositionAroundAboveGroundWithDistance(this Vector3 center,
			float radius,
			LayerMask layerMask,
			float distance,
			float prefabHeight = 1f,
			Transform transform = null,
			int numberOfTries = 10)
		{
			if (distance > radius)
			{
				Debug.LogError("Distance must be inferior to radius");
				return Vector3.zero;
			}
			var tries = 0;
			// While we didn't find a free position
			while (tries < numberOfTries)
			{
				// We pick a random position around above ground
				var newPos = center + Random.insideUnitSphere * radius;
				// center.y += 1000; // Security check, AboveGround check below first
				newPos = newPos.PositionAboveGround(prefabHeight, transform);
				if (newPos == Vector3.zero) // Outside map
				{
					tries++;
					continue;
				}
				// Then we check if this spot is free (from the given layer)
				var hitColliders = Physics.OverlapSphere(newPos, distance, layerMask);

				// If no objects of the same layer is detected, this spot is free, return
				if (hitColliders.Length == 0) return newPos;

				tries++;
			}

			return Vector3.zero;
		}
	}
}
